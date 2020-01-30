package smbios

var _ProcessorFamily = _MapDictionary{
	0x01:  "Other",
	0x02:  "Unknown",
	0x03:  "8086",
	0x04:  "80286",
	0x05:  "Intel386™ processor",
	0x06:  "Intel486™ processor",
	0x07:  "8087",
	0x08:  "80287",
	0x09:  "80387",
	0x0A:  "80487",
	0x0B:  "Intel® Pentium® processor",
	0x0C:  "Pentium® Pro processor",
	0x0D:  "Pentium® II processor",
	0x0E:  "Pentium® processor with MMX™ technology",
	0x0F:  "Intel® Celeron® processor",
	0x10:  "Pentium® II Xeon™ processor",
	0x11:  "Pentium® III processor",
	0x12:  "M1 Family",
	0x13:  "M2 Family",
	0x14:  "Intel® Celeron® M processor",
	0x15:  "Intel® Pentium® 4 HT processor",
	0x18:  "AMD Duron™ Processor Family",
	0x19:  "K5 Family",
	0x1A:  "K6 Family",
	0x1B:  "K6-2",
	0x1C:  "K6-3",
	0x1D:  "AMD Athlon™ Processor Family",
	0x1E:  "AMD29000 Family",
	0x1F:  "K6-2+",
	0x20:  "Power PC Family",
	0x21:  "Power PC 601",
	0x22:  "Power PC 603",
	0x23:  "Power PC 603+",
	0x24:  "Power PC 604",
	0x25:  "Power PC 620",
	0x26:  "Power PC x704",
	0x27:  "Power PC 750",
	0x28:  "Intel® Core™ Duo processor",
	0x29:  "Intel® Core™ Duo mobile processor",
	0x2A:  "Intel® Core™ Solo mobile processor",
	0x2B:  "Intel® Atom™ processor",
	0x2C:  "Intel® Core™ M processor",
	0x2D:  "Intel(R) Core(TM) m3 processor",
	0x2E:  "Intel(R) Core(TM) m5 processor",
	0x2F:  "Intel(R) Core(TM) m7 processor",
	0x30:  "Alpha Family [2]",
	0x31:  "Alpha 21064",
	0x32:  "Alpha 21066",
	0x33:  "Alpha 21164",
	0x34:  "Alpha 21164PC",
	0x35:  "Alpha 21164a",
	0x36:  "Alpha 21264",
	0x37:  "Alpha 21364",
	0x38:  "AMD Turion™ II Ultra Dual-Core Mobile M Processor Family",
	0x39:  "AMD Turion™ II Dual-Core Mobile M Processor Family",
	0x3A:  "AMD Athlon™ II Dual-Core M Processor Family",
	0x3B:  "AMD Opteron™ 6100 Series Processor",
	0x3C:  "AMD Opteron™ 4100 Series Processor",
	0x3D:  "AMD Opteron™ 6200 Series Processor",
	0x3E:  "AMD Opteron™ 4200 Series Processor",
	0x3F:  "AMD FX™ Series Processor",
	0x40:  "MIPS Family",
	0x41:  "MIPS R4000",
	0x42:  "MIPS R4200",
	0x43:  "MIPS R4400",
	0x44:  "MIPS R4600",
	0x45:  "MIPS R10000",
	0x46:  "AMD C-Series Processor",
	0x47:  "AMD E-Series Processor",
	0x48:  "AMD A-Series Processor",
	0x49:  "AMD G-Series Processor",
	0x4A:  "AMD Z-Series Processor",
	0x4B:  "AMD R-Series Processor",
	0x4C:  "AMD Opteron™ 4300 Series Processor",
	0x4D:  "AMD Opteron™ 6300 Series Processor",
	0x4E:  "AMD Opteron™ 3300 Series Processor",
	0x4F:  "AMD FirePro™ Series Processor",
	0x50:  "SPARC Family",
	0x51:  "SuperSPARC",
	0x52:  "microSPARC II",
	0x53:  "microSPARC IIep",
	0x54:  "UltraSPARC",
	0x55:  "UltraSPARC II",
	0x56:  "UltraSPARC Iii",
	0x57:  "UltraSPARC III",
	0x58:  "UltraSPARC IIIi",
	0x60:  "68040 Family",
	0x61:  "68xxx",
	0x62:  "68000",
	0x63:  "68010",
	0x64:  "68020",
	0x65:  "68030",
	0x66:  "AMD Athlon(TM) X4 Quad-Core Processor Family",
	0x67:  "AMD Opteron(TM) X1000 Series Processor",
	0x68:  "AMD Opteron(TM) X2000 Series APU",
	0x69:  "AMD Opteron(TM) A-Series Processor",
	0x6A:  "AMD Opteron(TM) X3000 Series APU",
	0x6B:  "AMD Zen Processor Family",
	0x70:  "Hobbit Family",
	0x78:  "Crusoe™ TM5000 Family",
	0x79:  "Crusoe™ TM3000 Family",
	0x7A:  "Efficeon™ TM8000 Family",
	0x80:  "Weitek",
	0x81:  "Available for assignment",
	0x82:  "Itanium™ processor",
	0x83:  "AMD Athlon™ 64 Processor Family",
	0x84:  "AMD Opteron™ Processor Family",
	0x85:  "AMD Sempron™ Processor Family",
	0x86:  "AMD Turion™ 64 Mobile Technology",
	0x87:  "Dual-Core AMD Opteron™ Processor Family",
	0x88:  "AMD Athlon™ 64 X2 Dual-Core Processor Family",
	0x89:  "AMD Turion™ 64 X2 Mobile Technology",
	0x8A:  "Quad-Core AMD Opteron™ Processor Family",
	0x8B:  "Third-Generation AMD Opteron™ Processor Family",
	0x8C:  "AMD Phenom™ FX Quad-Core Processor Family",
	0x8D:  "AMD Phenom™ X4 Quad-Core Processor Family",
	0x8E:  "AMD Phenom™ X2 Dual-Core Processor Family",
	0x8F:  "AMD Athlon™ X2 Dual-Core Processor Family",
	0x90:  "PA-RISC Family",
	0x91:  "PA-RISC 8500",
	0x92:  "PA-RISC 8000",
	0x93:  "PA-RISC 7300LC",
	0x94:  "PA-RISC 7200",
	0x95:  "PA-RISC 7100LC",
	0x96:  "PA-RISC 7100",
	0xA0:  "V30 Family",
	0xA1:  "Quad-Core Intel® Xeon® processor 3200 Series",
	0xA2:  "Dual-Core Intel® Xeon® processor 3000 Series",
	0xA3:  "Quad-Core Intel® Xeon® processor 5300 Series",
	0xA4:  "Dual-Core Intel® Xeon® processor 5100 Series",
	0xA5:  "Dual-Core Intel® Xeon® processor 5000 Series",
	0xA6:  "Dual-Core Intel® Xeon® processor LV",
	0xA7:  "Dual-Core Intel® Xeon® processor ULV",
	0xA8:  "Dual-Core Intel® Xeon® processor 7100 Series",
	0xA9:  "Quad-Core Intel® Xeon® processor 5400 Series",
	0xAA:  "Quad-Core Intel® Xeon® processor",
	0xAB:  "Dual-Core Intel® Xeon® processor 5200 Series",
	0xAC:  "Dual-Core Intel® Xeon® processor 7200 Series",
	0xAD:  "Quad-Core Intel® Xeon® processor 7300 Series",
	0xAE:  "Quad-Core Intel® Xeon® processor 7400 Series",
	0xAF:  "Multi-Core Intel® Xeon® processor 7400 Series",
	0xB0:  "Pentium® III Xeon™ processor",
	0xB1:  "Pentium® III Processor with Intel® SpeedStep™ Technology",
	0xB2:  "Pentium® 4 Processor",
	0xB3:  "Intel® Xeon® processor",
	0xB4:  "AS400 Family",
	0xB5:  "Intel® Xeon™ processor MP",
	0xB6:  "AMD Athlon™ XP Processor Family",
	0xB7:  "AMD Athlon™ MP Processor Family",
	0xB8:  "Intel® Itanium® 2 processor",
	0xB9:  "Intel® Pentium® M processor",
	0xBA:  "Intel® Celeron® D processor",
	0xBB:  "Intel® Pentium® D processor",
	0xBC:  "Intel® Pentium® Processor Extreme Edition",
	0xBD:  "Intel® Core™ Solo Processor",
	0xBE:  "Intel® Core™ 2 Processor", // See 7.5.2 Table 23 notes.
	0xBF:  "Intel® Core™ 2 Duo Processor",
	0xC0:  "Intel® Core™ 2 Solo processor",
	0xC1:  "Intel® Core™ 2 Extreme processor",
	0xC2:  "Intel® Core™ 2 Quad processor",
	0xC3:  "Intel® Core™ 2 Extreme mobile processor",
	0xC4:  "Intel® Core™ 2 Duo mobile processor",
	0xC5:  "Intel® Core™ 2 Solo mobile processor",
	0xC6:  "Intel® Core™ i7 processor",
	0xC7:  "Dual-Core Intel® Celeron® processor",
	0xC8:  "IBM390 Family",
	0xC9:  "G4",
	0xCA:  "G5",
	0xCB:  "ESA/390 G6",
	0xCC:  "z/Architecture base",
	0xCD:  "Intel® Core™ i5 processor",
	0xCE:  "Intel® Core™ i3 processor",
	0xCF:  "Intel® Core™ i9 processor",
	0xD2:  "VIA C7™-M Processor Family",
	0xD3:  "VIA C7™-D Processor Family",
	0xD4:  "VIA C7™ Processor Family",
	0xD5:  "VIA Eden™ Processor Family",
	0xD6:  "Multi-Core Intel® Xeon® processor",
	0xD7:  "Dual-Core Intel® Xeon® processor 3xxx Series",
	0xD8:  "Quad-Core Intel® Xeon® processor 3xxx Series",
	0xD9:  "VIA Nano™ Processor Family",
	0xDA:  "Dual-Core Intel® Xeon® processor 5xxx Series",
	0xDB:  "Quad-Core Intel® Xeon® processor 5xxx Series",
	0xDD:  "Dual-Core Intel® Xeon® processor 7xxx Series",
	0xDE:  "Quad-Core Intel® Xeon® processor 7xxx Series",
	0xDF:  "Multi-Core Intel® Xeon® processor 7xxx Series",
	0xE0:  "Multi-Core Intel® Xeon® processor 3400 Series",
	0xE4:  "AMD Opteron™ 3000 Series Processor",
	0xE5:  "AMD Sempron™ II Processor",
	0xE6:  "Embedded AMD Opteron™ Quad-Core Processor Family",
	0xE7:  "AMD Phenom™ Triple-Core Processor Family",
	0xE8:  "AMD Turion™ Ultra Dual-Core Mobile Processor Family",
	0xE9:  "AMD Turion™ Dual-Core Mobile Processor Family",
	0xEA:  "AMD Athlon™ Dual-Core Processor Family",
	0xEB:  "AMD Sempron™ SI Processor Family",
	0xEC:  "AMD Phenom™ II Processor Family",
	0xED:  "AMD Athlon™ II Processor Family",
	0xEE:  "Six-Core AMD Opteron™ Processor Family",
	0xEF:  "AMD Sempron™ M Processor Family",
	0xFA:  "i860",
	0xFB:  "i960",
	0x100: "ARMv7",
	0x101: "ARMv8",
	0x104: "SH-3",
	0x105: "SH-4",
	0x118: "ARM",
	0x119: "StrongARM",
	0x12C: "6x86",
	0x12D: "MediaGX",
	0x12E: "MII",
	0x140: "WinChip",
	0x15E: "DSP",
	0x1F4: "Video Processor",
	0x200: "RISC-V RV32",
	0x201: "RISC-V RV64",
	0x202: "RISC-V RV128",
}

// ProcessorFamily represents a processor family field.
type ProcessorFamily uint16

func (pf ProcessorFamily) String() string {
	return _ProcessorFamily.str(uint16(pf))
}
