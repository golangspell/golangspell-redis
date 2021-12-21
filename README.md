# golangspell-redis
Plugin with the Golang Spell commands for adding Redis capabilities to the project

## Golang Spell
The Core project contains the core commands (and the respective templates) of the platform [Golang Spell](https://github.com/golangspell/golangspell).

![alt text](https://golangspell.com/golangspell/blob/master/img/gopher_spell.png?raw=true)

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=cover.out

go tool cover -html=cover.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```

## Install
To install the golangspell-redis spell use the command

```sh
golangspell addspell github.com/golangspell/golangspell-redis golangspell-redis
```

## Update
To update the golangspell-redis version use the command

```sh
golangspell updatespell github.com/golangspell/golangspell-redis golangspell-redis
```

## Usage
Go to the project root and run the command:
```sh
golangspell redisinit
```

The command will add the Redis infrastructure files:

* domain/cache.go: contains the Cache Interface with the cache component specification and the GetCache() function, which returns the cache component available in the application context

* domain/lock.go: contains the Locker Interface containing the distributed Locker component specification, the Lock handle interface specification and the GetLocker() function, which returns the Locker component available in the application context

* gateway/redis/redis_client.go: contains the Redis cache implementation

* gateway/redis/redis_lock.go: contains the Redis lock implementation

For using the Locker feature follow this example:

```sh
	logger := config.GetLogger
	defer logger().Sync()

    lockIdentifier := "MyLockIdentifier"
    lock, err := domain.GetLocker().ObtainLock(lockIdentifier, 900*time.Millisecond)
    if err != nil {
		msg := fmt.Sprintf("Could not obtain lock for key %s. Message: %s\n", lockIdentifier, err.Error())
		logger().Error(msg)
		return err
	}

    //TODO: Add here the logic to be run protected from concurrent changes during your specified lock TTL

    _ = lock.Release()
```

For adding a value to the cache follow this example (in the example the TTL is 10 hours):

```sh
	logger := config.GetLogger
	defer logger().Sync()

    cacheKey := "MyCacheKey"
    cachedValue := "MyCachedValue"
    err := domain.GetCache().Set(cacheKey, cachedValue, 10*time.Hour)
    if err != nil {
		msg := fmt.Sprintf("Could not set a value in the cache. Message: %s\n", err.Error())
		logger().Error(msg)
		return err
	}
```

For getting a value from the cache follow this example:

```sh
	logger := config.GetLogger
	defer logger().Sync()

    cacheKey := "MyCacheKey"
    value, err := domain.GetCache().Get(cacheKey)
    if err != nil {
		msg := fmt.Sprintf("Could not set a value in the cache. Message: %s\n", err.Error())
		logger().Error(msg)
		return err
	}

    logger().Infof("My cache value is: %s", value)
```

<p align="center">
    <img src="https://github.com/golangspell/golangspell/blob/master/img/gopher_spell.png" width="350" alt="Golang Spell logo"/>
</p>
