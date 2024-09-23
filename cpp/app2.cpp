#include <windows.h>
#include <tlhelp32.h>
#include <iostream>
#include <fstream>
#include <thread>
#include <chrono>

// Function to list all processes and write to file
void ListProcesses() {
    // Open the file to write the process list
    std::ofstream outputFile("file.txt");
    if (!outputFile.is_open()) {
        std::cerr << "Failed to open file for writing." << std::endl;
        return;
    }

    // Take a snapshot of the current processes
    HANDLE hSnapshot = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
    if (hSnapshot == INVALID_HANDLE_VALUE) {
        std::cerr << "Failed to take snapshot of processes." << std::endl;
        outputFile.close();
        return;
    }

    PROCESSENTRY32 pe;
    pe.dwSize = sizeof(PROCESSENTRY32);

    // Retrieve information about the first process
    if (!Process32First(hSnapshot, &pe)) {
        std::cerr << "Failed to get process information." << std::endl;
        CloseHandle(hSnapshot);
        outputFile.close();
        return;
    }

    // Write header to file
    // using </br> because we buffer it to html file and load it to browser
    outputFile << "PID\t\tProcess Name </br>" << std::endl;
    outputFile << "------------------------------ </br>" << std::endl;

    // List all processes and write to file
    do {
        outputFile << pe.th32ProcessID << "\t\t" << pe.szExeFile << "</br>" << std::endl;
    } while (Process32Next(hSnapshot, &pe));

    CloseHandle(hSnapshot);
    outputFile.close();
}

int main() {
    while (true) {
        // Your code here
        std::cout << "Pulling list from OS.." << std::endl;

        std::this_thread::sleep_for(std::chrono::seconds(5));
        ListProcesses();
    }
    return 0;
}
