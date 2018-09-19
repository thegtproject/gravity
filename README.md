# Gravity
3D OpenGL engine for Go

## Gravity Development Tools
For consistent tooling purposes across platforms all common Gravity development tooling is being put into ./developing/tools as go source files. You can run these tools by using ```go run ./developing/tools <command>```

Currently available GDT commands:
⋅⋅⋅apitrace   -  trace full opengl state and opens qapitrace gui (requires apitrace)
⋅⋅⋅depgraph   -  graphs package dependency relationships of Gravity (requires graphiz)
