# Gravity
3D OpenGL engine for Go

## Building
Under the ```developing``` directory is the project code I use to test as I work on the engine itself. I keep it always buildable when commiting. To try it out, run the command ```gdt rundevapp``` or ```./gdt.bash rundevapp```


## Gravity Development Tools
For consistent tooling purposes across platforms all common Gravity development tooling is being put into ./developing/tools as go source files. You can run these tools by using ```go run ./developing/tools <command>```

Currently available GDT commands:  
   **apitrace**   -  trace full opengl state and opens qapitrace gui (requires apitrace)  
   **depgraph**   -  graphs package dependency relationships of Gravity (requires graphiz)  
   **genversion** -  generate the version.go source file for Gravity (requires git)
   **rundevapp**  -  go run's the developing/experimenting program under ./developing

## VMWare
If you are trying to test this on linux on vmware you will have to modify these values inside initGLFW() of gravity.go:
```
    glfw.WindowHint(glfw.ContextVersionMajor, 3)
    glfw.WindowHint(glfw.ContextVersionMinor, 3)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
```

This will work until I start implementing opengl 4.3 features in the future.

## Progress Screenshots

(latest to oldest):

![](doc/images/gravityss4.png)
---

![](doc/images/gravityss3.png)
---

![](doc/images/gravityss2.png)
---

![](doc/images/gravityss1.png)
---

## Progress Tracking
I'm still a novice programmer, so I thought it'd be fun to do additional things to track my progress on this project as naturally I'll learn to get better over time. For instance, every major refactor/restructure I do on the code, I create a graph image of my dependency relationships. You can already see it getting better over time:

(latest to oldest)

![](doc/images/depgraph9.png)
---

![](doc/images/depgraph8.png)
---

![](doc/images/depgraph7.png)
---

![](doc/images/depgraph6.png)
---

![](doc/images/depgraph5.png)
---

![](doc/images/depgraph4.png)
---

![](doc/images/depgraph3.png)
---

![](doc/images/depgraph2.png)
---

![](doc/images/depgraph1.png)
---
