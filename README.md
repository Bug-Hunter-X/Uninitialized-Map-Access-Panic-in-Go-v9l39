# Uninitialized Map Access Panic in Go

This repository demonstrates a common error in Go: accessing a key in an uninitialized map.  Attempting to read from an uninitialized map will result in a runtime panic, halting program execution.

**The Problem:**

Go maps are not automatically initialized.  If you try to access a key in a map that hasn't been created, a panic occurs.  This is different from languages where uninitialized map access might return a default value (e.g., null or 0).

**The Solution:**

Always explicitly initialize your maps before accessing their keys.  The example provided shows how to correctly initialize a map to avoid this issue.  Use `make()` to create a map with its initial capacity.  Checking for the key's existence before access is also an effective strategy to prevent the panic.
