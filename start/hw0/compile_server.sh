#!/bin/bash

rm -f server

if [ -f "server.c" ]; then
    echo "Компиляция C сервера..."
    gcc -o server server.c
elif [ -f "server.cpp" ]; then
    echo "Компиляция C++ сервера..."
    g++ -o server server.cpp
elif [ -f "server.go" ]; then
    echo "Компиляция Go сервера"
    go build -o server server.go
elif [ -f "server.py" ]; then
    echo "Python не требует компиляции"
    exit 0
elif [ -f "server.js" ]; then
    echo "JavaScript не требует компиляции"
    exit 0
elif [ -f "server.java" ]; then
    echo "Компиляция Java сервера..."
    javac server.java
else
    echo "Не найден файл сервера для компиляции"
    exit 1
fi

echo "Компиляция сервера завершена"