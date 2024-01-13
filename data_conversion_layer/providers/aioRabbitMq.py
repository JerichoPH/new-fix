import asyncio
from fastapi import FastAPI
import aio_pika

app = FastAPI()


class AioRabbitMq:
    _connection = None
    _channel = None
    _queue = None
    _queue_name = None

    async def connection(self, queue_name: str = "default-queue"):
        """
        创建链接
        :return:
        """
        if self._connection is None:
            self._connection = await aio_pika.connect_robust(
                host="127.0.0.1",
                virtualhost="default-vhost",
                login="admin",
                password="123123",
                port=5672,
            )
            self._channel = await self._connection.channel()
            self._queue = await self._channel.declare_queue(queue_name)

        return self

    async def send_message(self, message: str):
        """
        发送消息到RabbitMQ
        :param message: 消息内容
        :return:
        """
        await self.connection()
        async with self._channel.sender(self._queue.name) as sender:
            await sender.send(aio_pika.Message(body=message.encode()))

    async def consume_messages(self):
        """
        消费消息
        :return:
        """
        if self._queue is not None:
            async with self._queue.iterator() as queue_iter:
                async for message in queue_iter:
                    async with message.process():
                        received_message = message.body.decode()
                        # 处理消息...
                        print(f"Received message: {received_message}")

    async def close(self):
        """
        关闭链接
        :return:
        """
        self._channel.close()
        self._connection.close()
