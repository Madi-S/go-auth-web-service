FROM python:3.10-slim

RUN pip install uv

WORKDIR /app

COPY pyproject.toml uv.lock ./

RUN uv sync --frozen --no-install-project --no-dev

COPY . .

EXPOSE 8000

CMD ["uv", "run","uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]