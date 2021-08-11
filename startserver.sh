cd ~/server
PORT=8081 ./streepjes &
serve -s ./dist &
sleep 5
firefox --new-instance http://localhost:5000/login

