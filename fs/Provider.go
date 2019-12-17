package fs

import (
	"errors"
	"io"
	"os"
	"time"
)

// A FileSystem component provides a means of selecting an underlying file system.

// Provider provides access to one or more file systems. This should be implemented
// by a vendor-specific implementation.
type Provider interface {
	// GetDefault selects the default file system provided by the provider.
	GetDefault() Fs

	// Get allows another type of file system to be created or selected.
	Get(fsType string) Fs

	// Supports indicates if the specified file system type is supported.
	Supports(fsType string) bool
}

// File represents a file in the filesystem.
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	Readdir(count int) ([]os.FileInfo, error)
	Readdirnames(n int) ([]string, error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	WriteString(s string) (ret int, err error)
}

// Fs defines the contract with a file system implementation.
// Any simulated or real filesystem should implement this interface.
type Fs interface {
	// Create creates a file in the filesystem, returning the file and an
	// error, if any happens.
	Create(name string) (File, error)

	// Mkdir creates a directory in the filesystem, return an error if any
	// happens.
	Mkdir(name string, perm os.FileMode) error

	// MkdirAll creates a directory path and all parents that does not exist
	// yet.
	MkdirAll(path string, perm os.FileMode) error

	// Open opens a file, returning it or an error, if any happens.
	Open(name string) (File, error)

	// OpenFile opens a file using the given flags and the given mode.
	OpenFile(name string, flag int, perm os.FileMode) (File, error)

	// Remove removes a file identified by name, returning an error, if any
	// happens.
	Remove(name string) error

	// RemoveAll removes a directory path and any children it contains. It
	// does not fail if the path does not exist (return nil).
	RemoveAll(path string) error

	// Rename renames a file.
	Rename(oldname, newname string) error

	// Stat returns a FileInfo describing the named file, or an error, if any
	// happens.
	Stat(name string) (os.FileInfo, error)

	// The name of this FileSystem
	Name() string

	//Chmod changes the mode of the named file to mode.
	Chmod(name string, mode os.FileMode) error

	//Chtimes changes the access and modification times of the named file
	Chtimes(name string, atime time.Time, mtime time.Time) error
}

var (
	// ErrFileClosed may occur after an attempt to close a file.
	ErrFileClosed = errors.New("File is closed")

	// ErrOutOfRange may occur when a file is out of range.
	ErrOutOfRange = errors.New("Out of range")

	// ErrTooLarge may occur when a file is too large.
	ErrTooLarge = errors.New("Too large")

	// ErrFileNotFound will a occur when a file cannot be found in the file system.
	ErrFileNotFound = os.ErrNotExist

	// ErrFileExists will occur when a file already exists.
	ErrFileExists = os.ErrExist

	// ErrDestinationExists will occur when a destination already exists.
	ErrDestinationExists = os.ErrExist

	// ErrUnknownFileSystem will occur when a named file system type is not supported.
	ErrUnknownFileSystem = errors.New(("Unknown file system type"))
)

var (
	// OsFileSystem is a symbolic name for the fault OS file system
	OsFileSystem = "osFs"

	// MemMapFileSystem is symbolic name for a memory mapped file system
	MemMapFileSystem = "mmFs"

	// SFTPFileSystem is a symbolic name for a Secure File Transfer Protocol file system
	SFTPFileSystem = "sftpFs"
)
