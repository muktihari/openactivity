export class ActivityFile {
  fileId: FileId | null = null
  sessions: Array<Session> | null = null
  records: Array<Record> | null = null
}

export class FileId {
  manufacturer: string | null = null
  product: number | null = null
  timeCreated: Date | null = null
}

export class Session {
  sport: string | null = null
  subSport: string | null = null
  totalMovingTime: number | null = null
  totalElapsedTime: number | null = null
  totalTimerTime: number | null = null
  totalDistance: number | null = null
  totalAscent: number | null = null
  totalDescent: number | null = null
  totalCycles: number | null = null
  totalCalories: number | null = null
  avgSpeed: number | null = null
  maxSpeed: number | null = null
  avgHeartRate: number | null = null
  maxHeartRate: number | null = null
  avgCadence: number | null = null
  maxCadence: number | null = null
  avgPower: number | null = null
  maxPower: number | null = null
  avgTemperature: number | null = null
  maxTemperature: number | null = null
  avgAltitude: number | null = null
  maxAltitude: number | null = null
}

export class Record {
  altitude: number | null = null
  cadence: number | null = null
  distance: number | null = null
  heartRate: number | null = null
  positionLat: number | null = null
  positionLong: number | null = null
  speed: number | null = null
  timestamp: Date | null = null
}
