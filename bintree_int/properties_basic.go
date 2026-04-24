package bintree_int

// IsEmpty liefert true zurück, wenn das Element leer ist.
//
// Ein Element gilt als leer, wenn es keine Kinder hat.
// D.h. wenn die Kind-Pointer Left und Right beide nil sind.
// Der Wert des Elements spielt dabei keine Rolle.
func (e *Element) IsEmpty() bool {
	// Hinweis:
	// Prüfen Sie, ob beide Kind-Pointer nil sind.

	// TODO
	return false
}

// IsLeaf liefert true zurück, wenn das Element ein Blatt ist.
//
// Ein Element gilt als Blatt, wenn beide Kinder leer sind.
func (e *Element) IsLeaf() bool {
	// Hinweis:
	// Nutzen Sie die IsEmpty-Methode, um zu prüfen, ob beide Kinder leer sind,
	// aber stellen Sie sicher, dass das Element e selbst nicht leer ist.

	// TODO
	return false
}

// Count zählt die Anzahl der Elemente im Baum, beginnend bei diesem Element.
// Dabei werden nur nicht-leere Elemente gezählt.
func (e *Element) Count() int {
	// Hinweis:
	// Prüfen Sie zunächst, ob das aktuelle Element leer ist. Wenn ja, geben Sie 0 zurück.
	// Ansonsten nutzen Sie Count rekursiv, um die Anzahl der Elemente in den linken
	// und rechten Teilbäumen zu zählen.

	// TODO
	return 0
}

// Height berechnet die Höhe des Baums.
// Ein leeres Element hat die Höhe 0, ein Blatt hat die Höhe 1, usw.
func (e *Element) Height() int {
	// Hinweis:
	// Prüfen Sie zunächst, ob das aktuelle Element leer ist. Wenn ja, geben Sie 0 zurück.
	// Ansonsten berechnen Sie die Höhe der linken und rechten Teilbäume rekursiv.
	// Die Höhe von e ist dann 1 plus die größere der beiden Höhen.

	// TODO
	return 0
}
