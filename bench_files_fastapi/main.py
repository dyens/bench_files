from fastapi import FastAPI
from fastapi.responses import FileResponse
from pathlib import Path


path = Path(__file__).parents[1].joinpath('my_file')

app = FastAPI()


@app.get("/")
async def root():
    return FileResponse(path)
