import Container from "../../components/layouts/Container";
import OrdersLayout from "../../components/layouts/OrdersLayout";

export default function Dashboard() {
  return (
    <main>
      <Container>
        <OrdersLayout />
      </Container>
    </main>
  );
}
