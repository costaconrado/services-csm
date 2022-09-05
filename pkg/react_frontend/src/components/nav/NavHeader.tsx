import * as React from 'react';
// eslint-disable-next-line @typescript-eslint/ban-ts-ignore
// @ts-ignore
import { Dropdown, DropdownItem, DropdownToggle, Title } from '@patternfly/react-core';
import { CaretDownIcon } from '@patternfly/react-icons';

import './NavHeader.scss';

export type NavHeaderProps = {
  onPerspectiveSelected: () => void;
};

type PerspectiveDropdownItemProps = {
  activePerspective: string;
  onClick: (perspective: string) => void;
};

const PerspectiveDropdownItem: React.FC<PerspectiveDropdownItemProps> = ({
  activePerspective,
  onClick,
}) => {
  
  return (
    <DropdownItem
      key={perspective.properties.id}
      onClick={(e: React.MouseEvent<HTMLLinkElement>) => {
        e.preventDefault();
        onClick(perspective.properties.id);
      }}
      isHovered={perspective.properties.id === activePerspective}
    >
      <Title headingLevel="h2" size="md" data-test-id="perspective-switcher-menu-option">
        <span className="oc-nav-header__icon">
          <React.Suspense fallback={<>&emsp;</>}>
            <LazyIcon />
          </React.Suspense>
        </span>
        {perspective.properties.name}
      </Title>
    </DropdownItem>
  );
};

const ClusterIcon: React.FC = () => <span className="co-m-resource-icon">C</span>;

const NavHeader: React.FC<NavHeaderProps> = ({ onPerspectiveSelected }) => {
  const [activePerspective, setActivePerspective] = useActivePerspective();
  const [isPerspectiveDropdownOpen, setPerspectiveDropdownOpen] = React.useState(false);


  const onPerspectiveSelect = React.useCallback(
    (perspective: string): void => {
      if (perspective !== activePerspective) {
        setActivePerspective(perspective);
      }
      setPerspectiveDropdownOpen(false);
      onPerspectiveSelected?.();
    },
    [activePerspective, onPerspectiveSelected, setActivePerspective],
  );

  return (
    <>
        <div className="oc-nav-header">
          <Dropdown
            isOpen={isPerspectiveDropdownOpen}
            toggle={
              <DropdownToggle onToggle={() => setPerspectiveDropdownOpen(!isPerspectiveDropdownOpen)}>
                <Title headingLevel="h2" size="md">
                    'All Clusters'
                </Title>
              </DropdownToggle>
            }
            dropdownItems={[
                <DropdownItem
                  key='sales'
                  onClick={() => {
                    onPerspectiveSelect('sales');
                    setPerspectiveDropdownOpen(false);
                  }}
                >
                  All Clusters
                </DropdownItem>,
              ]}
          />
        </div>
    </>
  );
};

export default NavHeader;
