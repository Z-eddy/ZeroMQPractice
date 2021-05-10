cmake_minimum_required(VERSION 3.5)

project(ZeroMQPractice)


set(ZeroMQ_DIR $ENV{ZeroMQ_DIR})
find_package(ZeroMQ REQUIRED)

SET(CMAKE_INSTALL_PREFIX ${CMAKE_BINARY_DIR}/install/${CMAKE_BUILD_TYPE})

if( MINGW )
    if( CMAKE_BUILD_TYPE STREQUAL "Debug" )
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++2a -O0 -g -Wall -Wunused-function -pthread")
        set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -std=c99 -g -pthread")
        message("MINGW - Debug mode")
    elseif( CMAKE_BUILD_TYPE STREQUAL "Release" ) # debug
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -O2 -pthread -Wunused-function")
        set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -std=c99 -O2 -pthread")
        message("MINGW - Release mode")
    endif() # Debug
    #
    if( PROJECT_ARCHITECTURE STREQUAL "x64" ) # x64
    set( CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m64")
    elseif(PROJECT_ARCHITECTURE STREQUAL "x86") # x86
    set( CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")
    endif() # x64
    #set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} " )
    set( CMAKE_EXECUTABLE_SUFFIX ".exe" )
endif() # MINGW

if( APPLE )
    if( CMAKE_BUILD_TYPE STREQUAL "Debug" )
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -O0 -g -Wall -Wunused-function -pthread")
        set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -std=c99 -g -pthread")
        message("CLANG - Debug mode")
    elseif( CMAKE_BUILD_TYPE STREQUAL "Release" ) # debug
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -O2 -pthread -Wunused-function")
        set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -std=c99 -O2 -pthread")
        message("CLANG - Release mode")
    endif() # Debug
    #
    if( PROJECT_ARCHITECTURE STREQUAL "x64" ) # x64
    set( CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m64")
    elseif(PROJECT_ARCHITECTURE STREQUAL "x86") # x86
    set( CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")
    endif() # x64
    #set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} " )
endif() # CLANG

# Microsoft Visual C++ Compilers
if(MSVC)
	#set path for QT
	set(CMAKE_PREFIX_PATH $ENV{QtDir})

    if( CMAKE_BUILD_TYPE STREQUAL "Debug" )
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
    elseif(CMAKE_BUILD_TYPE STREQUAL "Release")
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /fp:fast /Gy /Oi /Oy /O2 /Ot /Zi /EHsc ")
    endif()
    if(CMAKE_CL_64)
        set( PROJECT_ARCHITECTURE "x64" )
    else(CMAKE_CL_64)
        set( PROJECT_ARCHITECTURE "x86" )
    endif(CMAKE_CL_64)
    #
    ADD_DEFINITIONS(-D_CRT_SECURE_NO_WARNINGS)
    #
    set( CMAKE_EXECUTABLE_SUFFIX ".exe" )

    	include_directories(${CMAKE_SOURCE_DIR}/include)#以库形式存在的第三方头文件
	file(GLOB PlatformDlls ${CMAKE_SOURCE_DIR}/bin/*.dll)
else()
	set(CMAKE_MODULE_PATH "${CMAKE_SOURCE_DIR}/cmake;${CMAKE_MODULE_PATH}")
	find_package (glog REQUIRED)
	include_directories (${GLOG_INCLUDE_DIR})

	file(GLOB PlatformDlls ${CMAKE_SOURCE_DIR}/bin/*.so)
endif()
file(COPY ${PlatformDlls} DESTINATION ${CMAKE_BINARY_DIR}/bin)
file(COPY ${PlatformDlls} DESTINATION ${CMAKE_INSTALL_PREFIX}/bin)

if(CMAKE_CXX_COMPILER_ID  STREQUAL "Clang")
  #SET (CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fexceptions")
elseif(CMAKE_CXX_COMPILER_ID STREQUAL "MSVC")
  #SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /EHc-")
endif()

#增加输出目录
set(PluginDirectory ${CMAKE_BINARY_DIR}/bin/plugin)
file(MAKE_DIRECTORY ${PluginDirectory})
#设置通用的搜索文件
set(SourceFiles "./*.h" "./*.hpp" "./*.cpp" "./*.ui" "./*.qml" "./*.qrc" "./*.json" "./*.js" "./*.xlsx")
#设置通用的链接库
#set(CommonLib GlobalStruct glog QtXlsxWriter qt-material-widgets)
set(CommonLib GlobalStruct glog QtXlsxWriter)
#设置Driver的链接库
#set(AllDrivers BaseDrivers HalconDrivers VisaDrivers NormalDrivers)
#设置DeviceModels的链接库
#set(AllDeviceModels BaseDeviceModels HalconDeviceModels NormalDeviceModels)

#设置输出路径
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

#设定库文件后缀名
set(CMAKE_DEBUG_POSTFIX "d")

#复制依赖库
file(COPY ${ZeroMQ_LIBRARY} DESTINATION ${CMAKE_BINARY_DIR}/bin)
file(COPY ${ZeroMQ_LIBRARY} DESTINATION ${CMAKE_INSTALL_PREFIX}/bin)

#add_subdirectory(${CMAKE_SOURCE_DIR}/thirdparty)
add_subdirectory(src)
