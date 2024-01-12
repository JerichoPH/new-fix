import arrow
import pika

from utils.stdBusiness import StdBusniess


class RabbitMq:
    _connection = None
    _channel = None
    _queue_name = None

    def __init__(self, queue_name="default-queue"):
        self.connect_rabbitmq()
        self._queue_name = queue_name
        self._channel.queue_declare(queue=queue_name)

    def connect_rabbitmq(self):
        credentials = pika.PlainCredentials("admin", "123123")
        self._connection = pika.BlockingConnection(
            pika.ConnectionParameters(
                host="127.0.0.1",
                port="5672",
                virtual_host="default-vhost",
                credentials=credentials,
            )
        )
        self._channel = self._connection.channel()

    def send_message(self, message, exchange=""):
        self._channel.basic_publish(
            exchange=exchange, routing_key=self._queue_name, body=message
        )
        print(" [x] Sent %r" % message)


if __name__ == "__main__":
    rabbit_mq = RabbitMq()
    rabbit_mq.send_message(StdBusniess("ping", {"time": arrow().now()}).to_json_str())
