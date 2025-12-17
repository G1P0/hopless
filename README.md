# HOPLESS

**HOPLESS** — терминальная игра про сетевые правила доступа.

Deny by default. Last match wins. 

## Что есть сейчас

* узлы: `client`, `router`, `server`
* правила `ALLOW / DENY`
* поддержка портов (`0 = ANY`)
* принцип **last match wins**
* терминальный интерфейс
* одна миссия

## Пример миссии

```
ALLOW client -> server :80
DENY  client -> server :22
```

## Основные команды

* `show mission` — показать цель
* `show rules` — список правил
* `add rule` — добавить правило
* `ping` — проверить доступ (from, to, port)
* `check mission` — проверить выполнение

## Пример

Добавь:

```
client -> server port 80 allow
client -> server port 22 deny
```

Потом:

```
check mission
```

## Архитектура

```
domain/  — сущности
engine/  — логика доступа
ui/      — CLI
cmd/     — entrypoint
```

---

## Что дальше

* hop-by-hop маршрутизация
* протоколы
* scan / trace
* шум и ограничения
* процедурные миссии
