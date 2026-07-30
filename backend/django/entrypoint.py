import os
import time
import subprocess
import sys

import psycopg2


def wait_for_postgres() -> None:
    host = os.environ.get("POSTGRES_HOST", "postgres")
    port = int(os.environ.get("POSTGRES_PORT", "5432"))
    user = os.environ.get("POSTGRES_USER", "shopino")
    password = os.environ.get("POSTGRES_PASSWORD", "shopino123")
    dbname = os.environ.get("POSTGRES_DB", "shopino_db")

    for i in range(30):
        try:
            conn = psycopg2.connect(
                host=host, port=port, user=user, password=password, dbname=dbname
            )
            conn.close()
            print("Postgres is ready")
            return
        except Exception as exc:
            print(f"Waiting ({i + 1}/30): {exc}")
            time.sleep(2)
    raise SystemExit("Postgres not ready")


def main() -> None:
    wait_for_postgres()
    subprocess.check_call([sys.executable, "manage.py", "migrate", "--noinput"])
    os.execvp(
    "gunicorn",
    [
        "gunicorn",
        "shopino.wsgi:application",
        "--bind",
        "0.0.0.0:8000",
        "--workers",
        "2",
    ],
)


if __name__ == "__main__":
    main()
