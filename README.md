[![Build Status](https://travis-ci.org/bhowmikp/sudoku-solver.svg?branch=master)](https://travis-ci.org/bhowmikp/sudoku-solver)

# sudoku-solver
Run app:
```sh
./build-and-run.sh
```

Delete container:
```sh
./delete-container.sh
```

Code inside container:
```sh
sudo docker run -it golang-app /bin/bash
```

Deploy to Heroku:
```sh
heroku container:push web
heroku container:release web
```
