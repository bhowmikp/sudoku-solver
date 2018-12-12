[![Build Status](https://travis-ci.org/bhowmikp/sudoku-solver.svg?branch=master)](https://travis-ci.org/bhowmikp/sudoku-solver)

# sudoku-solver
Run app:
```sh
sudo docker build -t golang-app .
sudo docker run -it --rm --name running-app golang-app
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
