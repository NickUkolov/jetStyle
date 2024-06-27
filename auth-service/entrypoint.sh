#!/bin/bash
exec uvicorn --host=0.0.0.0 --port="$AUTH_SERVICE_PORT" --log-level debug main:app