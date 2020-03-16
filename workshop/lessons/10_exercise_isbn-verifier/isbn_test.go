package main

import "testing"

func TestIsValidISBNShouldPassForValidISBN(t *testing.T) {
	isbn := "3-598-21508-8"
	if !IsValidISBN(isbn) {
		t.Errorf("Expected ISBN %s to be valid", isbn)
	}
}

func TestIsValidISBNShouldPassForValidISBNWithNonIntegerCheckDigit(t *testing.T) {
	isbn := "3-598-21507-X"
	if !IsValidISBN(isbn) {
		t.Errorf("Expected ISBN %s to be valid", isbn)
	}
}

func TestIsValidISBNShouldFailForTooShortISBN(t *testing.T) {
	isbn := "3-59-21508-8"
	if IsValidISBN(isbn) {
		t.Errorf("Expected ISBN %s to be invalid (too short)", isbn)
	}
}

func TestIsValidISBNShouldFailForTooLongISBN(t *testing.T) {
	isbn := "39-598-21508-8"
	if IsValidISBN(isbn) {
		t.Errorf("Expected ISBN %s to be invalid (too long)", isbn)
	}
}

func TestIsValidISBNShouldFailForInvalidCheckDigit(t *testing.T) {
	isbn := "3-598-21508-9"
	if IsValidISBN(isbn) {
		t.Errorf("Expected ISBN %s to be invalid due to wrong check digit", isbn)
	}
}