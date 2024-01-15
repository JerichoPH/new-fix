import asyncio
from itertools import product
from fastapi import FastAPI
import aio_pika
from settings import setting

app_setting = setting.get("./app.ini")


app = FastAPI()


class AioRabbitMq:
    _conn = None
    _channel = None
    _production_queue_names = []
    _consume_queue_names = []
    _production_queues = {}
    _consume_queues = {}

    def __init__(self):
        self.generate_conn()

    async def generate_conn(self):
        self._production_queue_names = app_setting.get(
            "rabbit-mq-service", "production-queue-names"
        ).split(",")
        self._consume_queue_names = app_setting.get(
            "rabbit-mq-service", "consume-queue-names"
        ).split(",")

        if self._conn is None:
            self._conn = await aio_pika.connect_robust(
                host="127.0.0.1",
                virtualhost="default-vhost",
                login="admin",
                password="123123",
                port=5672,
            )
            self._channel = await self._conn.channel()

        return self

    async def close_conn(self):
        if self._conn is not None:
            await self._conn.close()

    async def generate_queues(self):
        # 创建队列（生产者）
        for producation_queue_name in self._production_queue_names:
            # self._production_queues[
            #     producation_queue_name
            # ] = await self._channel.declare_queue(producation_queue_name)
            await self._channel.declare_queue(producation_queue_name)

        # 创建队列（消费者）
        for consume_queue_name in self._consume_queue_names:
            # self._consume_queues[
            #     consume_queue_name
            # ] = await self._channel.declare_queue(consume_queue_name)
            await self._channel.declare_queue(consume_queue_name)

        return self

    @property
    async def get_conn(self):
        return self._conn

    @property
    async def get_channel(self):
        return self._channel

    async def send_message(self, queue_name: str, message: str):
        """
        发送消息到RabbitMQ
        :param message: 消息内容
        :return:
        """
        # 声明或获取队列
        queue = await self._channel.declare_queue(queue_name)

        # 创建并发布消息
        await self._channel.default_exchange.publish(
            aio_pika.Message(body=message.encode()), routing_key=queue.name
        )

    # async def on_message(message: aio_pika.IncomingMessage):
    async def on_message(self, message: aio_pika.IncomingMessage):
        print("开始监听")
        try:
            process = message.process()
            print(f"收到消息{message.body.decode()}")
            await process.ack()
            # async with message.process():
            #     body = message.body.decode()
            #     print(f"Received message: {body}")
                # 在这里处理消息
                # ...
                # await message.ack()
        except Exception as e:
            await message.reject(request=True)
            print(f"处理消息错误：{e}")

    async def consume_messages(self):
        """
        消费消息
        :return:
        """
        queue = await self._channel.declare_queue("dcl")
        await queue.consume(self.on_message)
        # for consume_queue_name in self._consume_queue_names:
        #     queue = await self._channel.declare_queue(consume_queue_name)
        #     await queue.consume(self.on_message)

        await asyncio.Future()
