import asyncio
from fastapi import FastAPI
from providers.aioRabbitMq import AioRabbitMq
from routers.auth import router as auth_router
from routers.temp import router as temp_router
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # 允许所有的源
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.on_event("startup")
async def startup_event():
    pass
    # aio_rabbit_mq = await AioRabbitMq().conn()
    # asyncio.create_task(await aio_rabbit_mq.consume_messages())


@app.on_event("shutdown")
async def shutdown_event():
    # 清理资源，关闭连接等操作...
    pass


app.include_router(auth_router, prefix="/api", tags=["auth"])
app.include_router(temp_router, prefix="/api", tags=["temp"])
