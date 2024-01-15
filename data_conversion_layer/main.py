import asyncio
import json
from fastapi import FastAPI
from providers.aioRabbitMq import AioRabbitMq
from routers.auth import router as auth_router
from routers.temp import router as temp_router
from fastapi.middleware.cors import CORSMiddleware
import aio_pika

app = FastAPI()
aio_rabbit_mq_task: asyncio.Task = None

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # 允许所有的源
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.on_event("startup")
async def startup_event():
    aio_rabbit_mq = await AioRabbitMq().generate_conn()
    asyncio.create_task(aio_rabbit_mq.consume_messages())


@app.on_event("shutdown")
async def shutdown_event():
    pass


app.include_router(auth_router, prefix="/api", tags=["auth"])
app.include_router(temp_router, prefix="/api", tags=["temp"])
