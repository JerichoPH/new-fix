from fastapi.routing import APIRouter
from providers.aioRabbitMq import AioRabbitMq
from utils.stdBusiness import StdBusniess
import arrow

url_prefix = "/temp"

router = APIRouter()


@router.get(f"{url_prefix}/sendToRabbitMq")
async def send_to_rabbit_mq():
    aio_rabbit_mq = AioRabbitMq()
    now = arrow.now().format('YYYY-MM-DD HH:mm:ss')
    print(now)
    aio_rabbit_mq.send_message(message=StdBusniess(
        business_type="ping", content={"time": now}
    ).to_json_str())
