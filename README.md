# fileserver

## POST

### /file/create

curl -s -X POST http://localhost:3000/file/create -d 'uuid=uuid0001&name=额测试解耦&desc=这是一个ss个&url=http://127.dkjf/?dfdk&hash=sdfakdfajsdfk'

### /file/:uuid

http://localhost:3000/file/uuid0001

## upload file

<form action="http://123.206.29.15:3000/upload" method="post" enctype="multipart/form-data">
    Name: <input type="text" name="name"><br>
    Email: <input type="email" name="email"><br>
    Files: <input type="file" name="file"><br><br>
    <input type="submit" value="Submit">
</form>

进入/home/ubuntu/app/go,上传的文件如果成功,则在这个目录中