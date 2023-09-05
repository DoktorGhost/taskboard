# taskboard

docker build -t taskboard .
docker run -d --name taskboard-container -p 5432:5432 taskboard