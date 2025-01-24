Hi..I am Anubhav Ranjan and This is a Windows File Transfer Assignment..

Exercise 1: Freestyle Job - Windows File Transfer Automation 
1. Objective: Automate file transfer between two directories on a Windows 
machine. 
2. Steps: 
o Create a Freestyle Job in Jenkins. 
o Configure the job to execute a Windows batch command: 
xcopy C:\source-directory C:\target-directory /E /I /H /Y 
o Add a post-build action to check the success of the transfer using a 
custom message. 
3. Task: Test the job by placing some files in the source directory and verifying they 
are transferred to the target directory.