FROM python:3.10-slim

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

COPY requirements.txt /tmp/requirements.txt

RUN pip install --upgrade pip && pip3 --no-cache-dir install --prefer-binary -r /tmp/requirements.txt && rm /tmp/requirements.txt

RUN addgroup --gid 1000 notroot && \
    useradd -u 1000 -g notroot notroot -m

RUN mkdir /opt/auth_service

COPY --chown=notroot:notroot . /opt/auth_service

USER notroot
WORKDIR /opt/auth_service
RUN chmod +x /opt/auth_service/entrypoint.sh

EXPOSE $AUTH_SERVICE_PORT

CMD ["./entrypoint.sh"]
