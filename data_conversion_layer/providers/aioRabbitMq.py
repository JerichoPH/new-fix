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
    _consumer_queue_names = []
    _consume_queues = {}

    async def generate_conn(self):
        self._consumer_queue_names = app_setting.get(
            "rabbit-mq-service", "consumer-queue-names"
        ).split(",")

        if self._conn is None:
            self._conn = await aio_pika.connect_robust(
                host=app_setting.get("rabbit-mq-service", "addr"),
                virtualhost=app_setting.get("rabbit-mq-service", "vhost"),
                login=app_setting.get("rabbit-mq-service", "username"),
                password=app_setting.get("rabbit-mq-service", "password"),
                port=5672,
            )
            self._channel = await self._conn.channel()

        return self

    async def close_conn(self):
        if self._conn is not None:
            await self._conn.close()

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

    async def on_message(self, message: aio_pika.IncomingMessage):
        """
        处理消息
        :param message: 消息内容
        :return:
        """
        try:
            async with message.process():
                body = message.body.decode()
                print(f"Received message: {body}")

                with open("./a.json", mode="w") as f:
                    f.write(body)
        except:
            await message.reject()

    async def consume_messages(self):
        """
        消费消息
        :return:
        """
        for queue_name in self._consumer_queue_names:
            queue = await self._channel.declare_queue(queue_name)
            await queue.consume(self.on_message)
