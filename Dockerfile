FROM python:3.12-slim

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    CONFIG_PATH=/app/config.cfg

WORKDIR /app

# Install dependencies in a separate layer so source-only changes rebuild quickly.
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# config.cfg is intentionally not copied into the image. Mount it at runtime.
COPY bot.py ./

CMD ["python", "bot.py"]
