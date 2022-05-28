#!/bin/bash
set -e

# Dump the SQL DB to better understand changes between commits
sqlite3 ./internal/data/db.sqlite3 <<EOF
.output ./internal/data/dump.sql
.dump
.exit
EOF