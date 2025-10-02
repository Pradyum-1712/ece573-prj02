/*
 * Copyright 2023 Matthew A. Titmus
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
    "log"
    "os"
)

func getenvOr(k, def string) string {
    v := os.Getenv(k)
    if v == "" {
        return def
    }
    return v
}

func initializeTransactionLog() {
    params := PostgresDbParams{
        host:     getenvOr("POSTGRES_HOST", "postgres"),
        dbName:   getenvOr("POSTGRES_DB", "kvs"),
        user:     getenvOr("POSTGRES_USER", "test"),
        password: os.Getenv("POSTGRES_PASSWORD"),
    }
    if params.password == "" {
        log.Fatal("POSTGRES_PASSWORD must be set")
    }
    var err error
    transact, err = NewPostgresTransactionLogger(params)
    if err != nil {
        log.Fatalf("init txn log: %v", err)
    }
}
