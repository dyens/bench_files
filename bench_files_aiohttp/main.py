"""Response files server."""

from aiohttp import web
from pathlib import Path

path = Path(__file__).parents[1].joinpath('my_file')
async def response_file(request):
    """Response file.

    :param request: request
    """
    return web.FileResponse(path=path)


app = web.Application()
app.add_routes([web.get('/', response_file)])
web.run_app(app)
