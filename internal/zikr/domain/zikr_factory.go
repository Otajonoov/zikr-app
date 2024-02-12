package domain

type ZikrFactory struct{}

func NewZikrFactory() ZikrFactory {
	return ZikrFactory{}
}

func (z *ZikrFactory) ParseToDomain(arabic, uzbek, pronounce string) *Zikr {
	return &Zikr{
		arabic:    arabic,
		uzbek:     uzbek,
		pronounce: pronounce,
	}
}

//func (z *ZikrFactory) ParseToDomainArray(rows *pgx.Rows) *Zikrs {
//	var result Zikrs
//	result.Zikr = make([]*Zikr, 0)
//	log.Println(" : ", rows)
//
//	for rows.Next() {
//		var zikr Zikr
//		err := rows.Scan(
//			&zikr.arabic,
//			&zikr.uzbek,
//			&zikr.pronounce,
//		)
//		if err != nil {
//			return nil
//		}
//
//		result.Zikr = append(result.Zikr, &zikr)
//	}
//
//	return &result
//}
