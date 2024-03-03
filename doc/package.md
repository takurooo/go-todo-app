# パッケージ構成図

- パッケージ内の構成要素の関係を図示したもの。  
- パッケージ間の関係は図示していない。
- 実装上のインターフェース名は構造体名と被っているので、図では名前が被らないようにインターフェース名の先頭にIを付けている。

## handler package

```mermaid
classDiagram

    IAddTaskService <.. AddTaskHandler
    IListTasksService <.. ListTaskHandler
    IRegisterUserService <.. RegisterUserHandler
    ILoginService <.. LoginHandler

    namespace handler {
        class AddTaskHandler {
            +Service   AddTaskService
            +Validator *validator.Validate
            +ServeHTTP(w http.ResponseWriter, r *http.Request)
        }

        class ListTaskHandler {
            +Service   ListTasksService
            +ServeHTTP(w http.ResponseWriter, r *http.Request)
        }

        class RegisterUserHandler {
            +Service   RegisterUserService
            +Validator *validator.Validate
            +ServeHTTP(w http.ResponseWriter, r *http.Request)
        }

        class LoginHandler {
            +Service   LoginService
            +Validator *validator.Validate
            +ServeHTTP(w http.ResponseWriter, r *http.Request)
        }

        class IAddTaskService {
            <<interface>>
            +AddTask(ctx context.Context, title string) (*entity.Task, error)
        }
        class IListTasksService {
            <<interface>>
            +ListTasks(ctx context.Context) (entity.Tasks, error)
        }
        class ILoginService {
            <<interface>>
            +Login(ctx context.Context, name, pw string) (string, error)
        }
        class IRegisterUserService {
            <<interface>>
            +RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
        }
    }
```

## service package

```mermaid
classDiagram

    ITaskAdder <.. AddTaskService
    ITaskLister <.. ListTasksService
    IUserGetter <.. LoginService
    ITokenGenerator <.. LoginService
    IUserRegister <.. RegisterUserService

    namespace service {
        class AddTaskService {
            +DB store.Execer
            +Repo TaskAdder
            +AddTask(ctx context.Context, title string) (*entity.Task, error)
        }
        class ListTasksService {
            +DB   store.Queryer
            +Repo TaskLister
            +ListTasks(ctx context.Context) (entity.Tasks, error)
        }
        class LoginService {
            +DB             store.Queryer
            +Repo           UserGetter
            +TokenGenerator TokenGenerator
            +Login(ctx context.Context, name, pw string) (string, error)
        }
        class RegisterUserService {
            +DB   store.Execer
            +Repo UserRegister
            +RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
        }

        class ITaskAdder {
            <<interface>>
            +AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
        }
        class ITaskLister {
            <<interface>>
            +ListTasks(ctx context.Context, db store.Queryer, id entity.UserID) (entity.Tasks, error)
        }
        class IUserGetter {
            <<interface>>
            +GetUser(ctx context.Context, db store.Queryer, name string) (*entity.User, error)
        }
        class ITokenGenerator {
            <<interface>>
            +GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
        }
        class IUserRegister {
            <<interface>>
            +RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
        }

    }
```

## store package

```mermaid
classDiagram

    IExecer <.. Repository
    IQueryer <.. Repository
    IPreparer <.. IQueryer

    namespace store {
        class KVS {
            +Cli *redis.Client
            +Save(ctx context.Context, key string, userID entity.UserID) error
            +Load(ctx context.Context, key string) (entity.UserID, error)
        }
        class Repository {
            Clocker clock.Clocker
            AddTask(ctx context.Context, db Execer, t *entity.Task) error
            ListTasks(ctx context.Context, db Queryer, id entity.UserID) (entity.Tasks, error)
            RegisterUser(ctx context.Context, db Execer, u *entity.User) error
            GetUser(ctx context.Context, db Queryer, name string) (*entity.User, error)
        }
        class TaskStore {
            +LastID entity.TaskID
            +Tasks  map[entity.TaskID]*entity.Task
            +Add(t *entity.Task) (entity.TaskID, error)
            +All() entity.Tasks
        }

        class IBeginner {
            <<interface>>
            BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
        }
        class IPreparer {
            <<interface>>
            PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
        }
        class IExecer {
            <<interface>>
            ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
            NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
        }
        class IQueryer {
            <<interface>>
            QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
            QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
            GetContext(ctx context.Context, dest any, query string, args ...any) error
            SelectContext(ctx context.Context, dest any, query string, args ...any) error
        }
    }
```

## auth package

```mermaid
classDiagram

    Store <.. JWTer

    namespace auth {
        class JWTer {
            +PrivateKey jwk.Key
            +PublicKey  jwk.Key
            +Store      Store
            +Clocker    clock.Clocker
            +GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
            +GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) 
            +FillContext(r *http.Request) (*http.Request, error)
        }

        class Store {
            <<interface>>
            Save(ctx context.Context, key string, userID entity.UserID) error
            Load(ctx context.Context, key string) (entity.UserID, error)
        }
    }
```

## clock package

```mermaid
classDiagram

    Clocker <|.. RealClocker
    Clocker <|.. FixedClocker

    namespace clock {
        class Clocker {
            <<interface>> 
            +Now() time.Time
        }
        class RealClocker {
            +Now() time.Time
        }
        class FixedClocker {
            +Now() time.Time
        }
    }
```
