package mdt

type AKSxx1603 struct {
	Socket bool
}

func (a *AKSxx1603) Functions() []string {
	return []string{
		"Schalten_1B",
		"Treppenlicht_1B",
		"Service_erforderlich_1B",
		"Schaltimpuls_1B",
		"Sperren_1B",
		"Zeit_bis_nächsten_Service_2By_4By",
		"Betriebsstunden_RM_2By_4By",
		"Treppenlicht_mit_Zeit_1By",
		"Vorwarnen_1B",
		"Betriebsstunden_Reset_1B",
		"Servicemeldung_Reset_1B",
		"Zwangsführung_2B",
		"Priorität_1B",
		"Szene_1By",
		"Status_1B",
		"Status_invertiert_1B",
		"Schwellwertschalter_1By_2By",
		"Logik_1_1B",
		"Logik_2_1B",
	}
}

func (a *AKSxx1603) Prefix() string {
	if a.Socket {
		return "S"
	}
	return "L"
}
