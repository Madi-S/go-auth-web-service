from fastapi import FastAPI
from src.routers.dummy import router as dummy_router

app = FastAPI()

app.include_router(dummy_router)


@app.get("/ping")
def ping():
    return "pong"


@app.get("/healthz")
def health_check():
    return {"status": "ok"}
