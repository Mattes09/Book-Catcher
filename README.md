# Book-Catcher
Program vyhľadá konkrétnu knižku podľa ISBN čísla, ktorý USER zadá do command linu.
Za použitia OPEN LIBRARY APIS a to konkrétne ISBN API, /AUTHORS API, /AUTHORS/WORKS API som sa dostal ya pomocou Unmrsahall k dátam autorov a ich prác.
Snažil som sa taktiež použiť YAML formát ako je vyžiadané v zadaní avšak API v tomto formáte nefungovalo pre AUTHORS/WORKS 
Zatiaľ som sa nedopracoval ku koverzii JSON na YAML 
Taktiež program funguje celkom spoľahlivo avšak našiel som niekoľko chýb kde nezobrazí meno autora - štruktúra JSONu nie je rovnaká pri všetkých knihách