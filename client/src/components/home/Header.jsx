import Container from "../layouts/Container";
import TrackingForm from "./TrackingForm";

export default function Header() {
  return (
    <section className="relative bg-[url('./assets/shipment.jpg')] bg-coverflex items-center justify-center bg-center mb-20">
      <Container>
        <div className="relative z-20 flex flex-col item-center justify-center py-28 text-center text-white">
          <h1 className="text-4xl mb-14">Shipping and logistics service provider</h1>
          <h5 className="text-xl mb-5">
            TRACK YOUR ORDER
          </h5>
         <TrackingForm />
        </div>
      </Container>

      <div className="bg-black/75 absolute inset-0 z-10"></div>
    </section>
  );
}
