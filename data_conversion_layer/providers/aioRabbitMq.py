import asyncio
from itertools import product
from fastapi import FastAPI
import aio_pika
from settings import setting

app_setting = setting.get("./app.ini")


app = FastAPI()

connection = None
channel = None
production_queue_names = app_setting.get(
    "rabbit-mq-service", "production-queue-names"
).split(",")
consume_queue_names = app_setting.get("rabbit-mq-service", "consume-queue-names").split(
    ","
)

production_queues: dict = {}
consume_queues: dict = {}


class AioRabbitMq:
    _conn = None
    _channel = None
    _production_queue_names = []
    _consume_queue_names = []
    _production_queues = {}
    _consume_queues = {}

    async def conn(self):
        self._production_queue_names = app_setting.get(
            "rabbit-mq-service", "production-queue-names"
        ).split(",")
        self._consume_queues = app_setting.get(
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

    async def generate_queues(self):
        # 创建队列（生产者）
        for producation_queue_name in self._production_queue_names:
            production_queues[
                producation_queue_name
            ] = await self._channel.declare_queue(producation_queue_name)

        # 创建队列（消费者）
        for consume_queue_name in self._consume_queue_names:
            consume_queues[consume_queue_name] = await self._channel.declare_queue(
                consume_queue_name
            )

        return self

    async def send_message(self, queue_name: str, message: str):
        """
        发送消息到RabbitMQ
        :param message: 消息内容
        :return:
        """
        async with self._channel.sender(production_queues[queue_name].name) as sender:
            await sender.send(aio_pika.Message(body=message.encode()))

    async def consume_messages(self):
        """
        消费消息
        :return:
        """
        for consume_queue in self._consume_queues:
            async with consume_queue.iterator() as queue_iter:
                async for message in queue_iter:
                    async with message.process():
                        received_message = message.body.decode()
                        # 处理消息...
                        print(f"Received message: {received_message}")
