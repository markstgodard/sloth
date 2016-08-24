# sloth (⊙ω⊙)
Simple app with a configurable delay.

## Usage
Push and start app with configurable delay duration
```
cf push sloth --no-start
cf set-env sloth DELAY_TIME "100ms"
cf start sloth
```

Example output
```
$ curl sloth.bosh-lite.com
 endpoint: 10.244.0.26:61012 slept for: 100ms
```
