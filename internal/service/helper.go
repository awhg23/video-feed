package service

func uniqueUint64(ids []uint64) []uint64 {
	seen := make(map[uint64]struct{}, len(ids))
	res := make([]uint64, 0, len(ids))

	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}

		seen[id] = struct{}{}
		res = append(res, id)
	}

	return res
}
