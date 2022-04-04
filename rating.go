package main

import "fmt"

type Playstyle struct {
	Infantry       float32
	Cavalry        float32
	Archers        float32
	CavalryArchers float32
	Monks          float32
	Siege          float32
	Water          float32
}

func (r *Playstyle) Versatility() float32 {
	return (r.Infantry + r.Cavalry + r.Archers + r.CavalryArchers + r.Monks + r.Siege + r.Water)
}

func GetPlaystyle(civilizationName string) (Playstyle, error) {
	ratings := map[string]Playstyle{}

	ratings["Aztecs"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Berbers"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Britons"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Bulgarians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.0, CavalryArchers: 1.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Burgundians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Burmese"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Byzantines"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Celts"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Chinese"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Cumans"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Ethiopians"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Franks"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Goths"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Huns"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Incas"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Indians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Italians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Japanese"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Khmer"] = Playstyle{Infantry: 0.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Koreans"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Lithuanians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Magyars"] = Playstyle{Infantry: 0.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Malay"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Malians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Mayans"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Mongols"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Persians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.0, Siege: 1.0, Water: 1.0}
	ratings["Portuguese"] = Playstyle{Infantry: 0.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Saracens"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Sicilians"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Slavs"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Spanish"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.0, CavalryArchers: 1.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Tatars"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.3, Siege: 1.0, Water: 1.0}
	ratings["Teutons"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 0.5, CavalryArchers: 0.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Turks"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 1.0, Siege: 1.0, Water: 1.0}
	ratings["Vietnamese"] = Playstyle{Infantry: 1.0, Cavalry: 1.0, Archers: 1.0, CavalryArchers: 1.0, Monks: 0.7, Siege: 1.0, Water: 1.0}
	ratings["Vikings"] = Playstyle{Infantry: 1.0, Cavalry: 0.0, Archers: 1.0, CavalryArchers: 0.0, Monks: 0.3, Siege: 1.0, Water: 1.0}

	if rating, ok := ratings[civilizationName]; ok {
		return rating, nil
	}

	return Playstyle{}, fmt.Errorf("unknown civilization: %s", civilizationName)
}
