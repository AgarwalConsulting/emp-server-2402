# ReSTful (Respresentational State Transfer)

HTTP Methods => Get, Put, Post, Delete, Options, ...

CRUD => Create, Read, Update, Destroy

## Employee Management Server (JSON API)

```
CRUD        Action          HTTP Method               URI                         Req Body                    Resp Body
---------------------------------------------------------------------------------------------------------------------------
Read        Index              GET                 /employees                        -                       [{...}, ...] - Done
Read        Show               GET                /employees/{id}                    -                        {...}
Create      Create            POST                 /employees                      {...}                      {id: ..., ...} - Done
Update      Update             PUT                /employees/{id}                  {...}                      - / {...}
Update      Update           PATCH                /employees/{id}                {some attrs}                 - / {...}
Destroy     Delete           DELETE               /employees/{id}                    -                        - / {...}
```
