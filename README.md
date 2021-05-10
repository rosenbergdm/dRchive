# dRchive

A program for long-term file backup and verification

## Plan

There are three main *locations* and one configuration directory:

1. The **SOURCE** directory.  This is essentially the working tree, etc. that is
    snapshotted.
2. The **IMAGE** directory.  This is a snapshot of the working tree at the time
    of the last backup.
3. The **ARCHIVE** directory.  The (possibly offsite, possibly compressed)
    version of the **IMAGE**.
4. The **CONFIG** directory.  This is the location of the configuration data
    and sqlite3 databases.

The plan for each interval backup/check step is:

1. The **SOURCE** directory is walked.  For each file:

- If file is in the sqlite3 db
  - If the file mtime is newer than that in the sqlite db:
    - This is an *update*: The file is copied to the **IMAGE** directory and
          md5sums are calculated to confirm that the file is copied correctly
    - The updated md5sum and mtime are written to the sqlite database and the
          lastactive field is set to the date of the backup
  - The file mtime matches the db:
    - This is a *verification*.  The file in the **IMAGE** directory and the
          file in the **SOURCE** directory have their md5sums verified as equal
    - The lastactive field is set to the date of the backup
- If the file is not in the **IMAGE** directory:
  - This is a *creation*
  - The directories below the file are created in the **IMAGE** directory
  - The file is copied into the directory
  - The md5 of the copied file is verified as the same as the source
  - A new entry is made in the database
- All files with lastactive earlier than the date of the backup are removed
      from the **IMAGE** directory

2. For each **ARCHIVE** directory:

- The **IMAGE** directory is sent to the **ARCHIVE** target (could be a
      backblaze backup, compressed file, rsync'd remote location)
- If possible, the **ARCHIVE** is verified as identical to the **IMAGE**
      (which will be dependent on the archive mode type)

3. A log report is created.  The sqlite3 database is backed up.  A file with the
    sqlite3 database hash is created as well.

## Architecture (language? Is this a good project for go?)

- It'd be good to make the source directory read-only for the time of the backup.
- Will have to have a way to handle files it can't read (permissions), files of
    wrong type (devices, pipes, etc), symlinks/hardlinks, and recovery from a
    crash

