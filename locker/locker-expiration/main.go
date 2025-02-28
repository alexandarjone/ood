package main

import (
	"fmt"
	"time"
)

// Package represents a package in the locker.
type Package struct {
	id      int
	expired bool
}

// Expire marks the package as expired.
func (p *Package) Expire() {
	p.expired = true
	fmt.Printf("Package %d has expired.\n", p.id)
}

// PackageExpirationListener listens for expiration timeout or cancellation.
type PackageExpirationListener struct {
	pkg      *Package
	duration time.Duration
	cancelCh chan struct{}
}

// NewPackageExpirationListener creates a new listener for a package.
func NewPackageExpirationListener(pkg *Package, duration time.Duration) *PackageExpirationListener {
	return &PackageExpirationListener{
		pkg:      pkg,
		duration: duration,
		cancelCh: make(chan struct{}),
	}
}

// Start begins the waiting process in a separate goroutine.
func (l *PackageExpirationListener) Start() {
	go func() {
		select {
		case <-time.After(l.duration):
			l.pkg.Expire()
		case <-l.cancelCh:
			fmt.Printf("Expiration listener for package %d cancelled.\n", l.pkg.id)
		}
	}()
}

// Cancel stops the expiration listener.
func (l *PackageExpirationListener) Cancel() {
	close(l.cancelCh)
}

func main() {
	// Create a package.
	pkg := &Package{id: 1}

	// Create an expiration listener for the package.
	// Using 5 seconds for demonstration; use 48*time.Hour in production.
	listener := NewPackageExpirationListener(pkg, 5*time.Second)
	fmt.Println("Package placed in locker. Waiting for expiration...")
	listener.Start()

	// Simulate picking up the package before it expires.
	time.Sleep(3 * time.Second)
	fmt.Println("Package picked up! Cancelling expiration listener.")
	listener.Cancel()

	// Wait to confirm that the package does not expire.
	time.Sleep(3 * time.Second)
	fmt.Printf("Final expired status for package %d: %v\n", pkg.id, pkg.expired)
}
