# **ISPmyAdmin**

ISPmyAdmin is a powerful and user-friendly tool designed to help Internet Service Providers (ISPs) manage their operations efficiently. It includes features such as a RADIUS server for authentication, an Admin CLI for managing users, and a GUI for enhanced usability. The project is built using Go and leverages modern libraries like GORM for database management and Fyne for the GUI.

---

## **Features**
1. **RADIUS Server**:
   - Handles authentication requests.
   - Validates credentials against a MySQL database.
   - Supports RADIUS clients like `radtest` and `NTRadPing`.

2. **Admin CLI**:
   - Add and manage users.
   - View all users in the database.
   - Execute Telnet commands for OLT (Optical Line Terminal) management.

3. **GUI**:
   - User-friendly interface for managing users.
   - View logs and manage RADIUS server operations.

4. **Database Integration**:
   - Uses MySQL for storing user credentials.
   - Automatically creates tables using GORM's `AutoMigrate`.

---

## **Installation**

### **1. Prerequisites**
- **Go**: Install Go from [golang.org](https://golang.org/).
- **XAMPP**: Install XAMPP for MySQL and Apache.
- **MySQL Database**: Ensure the MySQL service is running in XAMPP.

### **2. Clone the Repository**
```bash
git clone <repository-url>
cd ISPmyAdmin