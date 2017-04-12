call gengen_build.bat
@if errorlevel 1 goto failed
call gen_all.bat 
@if errorlevel 1 goto failed
revel run github.com/three-plus-three/web_example
@goto :eof

:failed
@exit /b -1