package tempconv

type meter int
type kilometer int

const (
    None         meter = 0
    OneKilometer meter = 0
)

func Meter(m meter) meter {
    return m
}

func KiloMeter(k kilometer) kilometer {
    return k
}

func MToK(m meter) kilometer {
    return kilometer(m / 1000)
}

func KToM(k kilometer) meter {
    return meter(k * 1000)
}
