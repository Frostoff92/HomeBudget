# 🏠 HomeBudget Tracker

**Pet-project:** приложение для учёта и прогнозирования домашних расходов.  
Создано для практики **DevOps**, **Python**, **SQL** и **CI/CD** — от проектирования базы данных до развёртывания и мониторинга.

---

## ⚙️ Технологический стек

| Область | Технологии |
|:--|:--|
| Backend | [FastAPI](https://fastapi.tiangolo.com/), [Python 3.11+](https://www.python.org/), [SQLAlchemy](https://www.sqlalchemy.org/), [Alembic](https://alembic.sqlalchemy.org/) |
| Database | [PostgreSQL](https://www.postgresql.org/) (в проде), [SQLite](https://www.sqlite.org/) (локально) |
| Frontend | [React](https://react.dev/), [Next.js](https://nextjs.org/), [Recharts](https://recharts.org/) |
| Infrastructure | [Docker](https://www.docker.com/), [docker-compose](https://docs.docker.com/compose/), [Helm](https://helm.sh/), [Kubernetes](https://kubernetes.io/) |
| CI/CD | [GitHub Actions](https://docs.github.com/en/actions) |
| Monitoring | [Prometheus](https://prometheus.io/) + [Grafana](https://grafana.com/) |
| Testing | pytest, HTTPX |

---

## 🎯 Цели проекта

- 💡 Реализовать полноценный DevOps-цикл: **от кода до продакшена**
- 📊 Отслеживать и анализировать **ежедневные расходы**
- 💰 Контролировать **ежемесячный бюджет по категориям**
- 🔮 Прогнозировать траты до конца месяца
- 📦 Экспортировать/импортировать покупки в CSV
- 🚨 Получать уведомления при перерасходе бюджета *(в планах)*

---

## 🧠 Основные возможности

- Добавление покупок: товар, количество, цена, магазин, категория  
- Расчёт общей суммы чека и автоматическое распределение по категориям  
- Дашборд расходов за день/неделю/месяц  
- Прогноз бюджета до конца месяца (по скользящей средней)  
- Экспорт в CSV и просмотр аналитики в Grafana  
- REST API со Swagger-документацией  

---

## 🏗️ Архитектура проекта

