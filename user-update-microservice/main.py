from fastapi import FastAPI
from app.routes.update_routes import router as update_router

app = FastAPI()

# Incluir las rutas en la aplicación
app.include_router(update_router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=4002)
