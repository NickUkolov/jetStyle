import uvicorn
import yaml
from fastapi import FastAPI, Response
from fastapi.openapi.utils import get_openapi

from routers.users import router as user_router

app = FastAPI()
app.include_router(user_router)


@app.get("/openapi_yaml", include_in_schema=True)
def get_openapi_yaml():
    openapi_schema = get_openapi(
        title="Auth Service API",
        version="1.0.0",
        description="API for user authentication and management",
        routes=app.routes,
    )
    yaml_schema = yaml.dump(openapi_schema)
    return Response(content=yaml_schema, media_type="application/x-yaml")


if __name__ == '__main__':
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)
