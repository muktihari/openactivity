export class ActivityFile {
  creator: Creator = new Creator()
  timezone: number = 0
  laps: Lap[] = []
  sessions: Session[] = []
  records: Record[] = []

  constructor(json?: any) {
    const casted = json as ActivityFile
    this.creator = casted?.creator
    this.timezone = casted?.timezone
    this.laps = casted?.laps
    this.sessions = casted?.sessions
    this.records = casted?.records
  }
}

export class Creator {
  name: string = 'Unknown'
  manufacturer: number = 0
  product: number = 0
  timeCreated: string | null = null
}

export class Session {
  sport: string | null = null
  totalMovingTime: number | null = null
  totalElapsedTime: number | null = null
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

export class Lap {
  timestamp: string | null = null
  totalMovingTime: number | null = null
  totalElapsedTime: number | null = null
  totalDistance: number | null = null
  totalAscent: number | null = null
  totalDescent: number | null = null
  totalCalories: number | null = null
  avgSpeed: number | null = null
  maxSpeed: number | null = null
  avgHeartRate: number | null = null
  maxHeartRate: number | null = null
  avgCadence: number | null = null
  maxCadence: number | null = null
  avgPower: number | null = null
  maxPower: number | null = null

  constructor(data?: any) {
    const casted = data as Lap
    this.timestamp = casted?.timestamp
    this.totalMovingTime = casted?.totalMovingTime
    this.totalElapsedTime = casted?.totalElapsedTime
    this.totalDistance = casted?.totalDistance
    this.totalAscent = casted?.totalAscent
    this.totalDescent = casted?.totalDescent
    this.totalCalories = casted?.totalCalories
    this.avgSpeed = casted?.avgSpeed
    this.maxSpeed = casted?.maxSpeed
    this.avgHeartRate = casted?.avgHeartRate ? casted?.avgHeartRate : null
    this.maxHeartRate = casted?.maxHeartRate ? casted?.maxHeartRate : null
    this.avgCadence = casted?.avgCadence
    this.maxCadence = casted?.maxCadence
    this.avgPower = casted?.avgPower ? casted?.avgPower : null
    this.maxPower = casted?.maxPower ? casted?.maxPower : null
  }
}

export class Record {
  timestamp: string | null = null
  positionLat: number | null = null
  positionLong: number | null = null
  distance: number | null = null
  speed: number | null = null
  altitude: number | null = null
  cadence: number | null = null
  heartRate: number | null = null
  power: number | null = null
  temperature: number | null = null
  grade: number = 0

  constructor(data?: any) {
    const casted = data as Record
    this.positionLat = casted?.positionLat
    this.positionLong = casted?.positionLong
    this.altitude = casted?.altitude
    this.cadence = casted?.cadence
    this.distance = casted?.distance
    this.heartRate = casted?.heartRate
    this.speed = casted?.speed
    this.timestamp = casted?.timestamp
    this.power = casted?.power
    this.temperature = casted?.temperature
    this.grade = casted?.grade
  }
}