function EventCard({ title, date, location }) {
  return (
    <article>
      <h3>{title}</h3>
      <p>Date: {date}</p>
      <p>Location: {location}</p>
    </article>
  )
}

export default EventCard