# CMakeList.txt : CMake project for ServerTest, include source and define
# project specific logic here.
#
cmake_minimum_required (VERSION 3.8)

project(ServerTest LANGUAGES CXX)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

file(GLOB SourceFiles ${SourceFiles})

add_executable(${PROJECT_NAME}
  ${SourceFiles}
)

INSTALL(TARGETS ${PROJECT_NAME}
RUNTIME DESTINATION bin
LIBRARY DESTINATION lib
ARCHIVE DESTINATION libstatic
)

target_link_libraries(${PROJECT_NAME} PRIVATE libzmq)
