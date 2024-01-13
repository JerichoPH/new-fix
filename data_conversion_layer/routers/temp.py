from fastapi.routing import APIRouter
from providers.aioRabbitMq import  AioRabbitMq
from utils.stdBusiness import StdBusniess
import arrow

from utils.stdResponse import StdResponse

url_prefix = "/temp"

router = APIRouter()


@router.get(f"{url_prefix}/sendToRabbitMq")
async def send_to_rabbit_mq():
    aio_rabbit_mq = await AioRabbitMq().conn()
    aio_rabbit_mq = await aio_rabbit_mq.generate_queues()
    now = arrow.now().format("YYYY-MM-DD HH:mm:ss")

    await aio_rabbit_mq.send_message(
        queue_name="new-fix-backend",
        message=StdBusniess(business_type="ping", content={"time": now}).to_json_str(),
    )
    return StdResponse().success("发送成功")
