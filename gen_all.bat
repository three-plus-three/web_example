call gengen.bat mvc -override=true -controller=App -projectPath=github.com/three-plus-three/web_example -customPath=/self -root specs  -view_tag=box
@if errorlevel 1 goto failed
call gengen.bat db -override=true -root specs -output=app/models
@if errorlevel 1 goto failed
call gengen.bat test_base -override=true -projectPath=github.com/three-plus-three/web_example -root specs 
@if errorlevel 1 goto failed

FOR %%i IN (app\controllers\*.go) DO goimports -w %%i
FOR %%i IN (app\models\*.go) DO goimports -w %%i
FOR %%i IN (tests\*.go) DO goimports -w %%i
:failed
@exit /b -1