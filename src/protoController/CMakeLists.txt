cmake_minimum_required(VERSION 3.14)

project(protoController LANGUAGES CXX)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

include($ENV{gRPC_DIR}/gRPCcommon.cmake)

######自动编译.proto######
# Proto file
get_filename_component(hw_proto "*.proto" ABSOLUTE)
get_filename_component(hw_proto_path "${hw_proto}" PATH)
file(MAKE_DIRECTORY ${CMAKE_BINARY_DIR}/protocCompiledFiles)
execute_process( COMMAND cmd /c
      protoc
      --cpp_out=${CMAKE_BINARY_DIR}/protocCompiledFiles
      -I ${hw_proto_path}
      ${hw_proto}
)
# Include generated *.pb.h files
include_directories(${CMAKE_BINARY_DIR}/protocCompiledFiles)
file(GLOB protocCompiledFiles ${CMAKE_BINARY_DIR}/protocCompiledFiles/*.h ${CMAKE_BINARY_DIR}/protocCompiledFiles/*.cc)


find_package(QT NAMES Qt6 Qt5 COMPONENTS Core REQUIRED)
find_package(Qt${QT_VERSION_MAJOR} COMPONENTS Core REQUIRED)

file(GLOB SourceFiles ${SourceFiles} *.proto)
add_library(${PROJECT_NAME}
    ${SourceFiles}
    ${protocCompiledFiles}
)

target_link_libraries(${PROJECT_NAME} PRIVATE ${gRPCAllLibs} Qt${QT_VERSION_MAJOR}::Core)